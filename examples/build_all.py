#!/usr/bin/env python3
import subprocess
from argparse import ArgumentParser
from pathlib import Path
from shutil import copytree, rmtree

import yaml
from jinja2 import Environment, FileSystemLoader


ROOT = Path(__file__).absolute().parent
env = Environment(
    loader=FileSystemLoader(ROOT),
    extensions=['jinja2_markdown.MarkdownExtension'],
)

parser = ArgumentParser()
parser.add_argument(
    '-o', '--output',
    default=str(ROOT.parent / 'public'),
    help='path to build output',
)


def make_index():
    with (ROOT / 'index.yml').open(encoding='utf8') as stream:
        data = yaml.safe_load(stream)
    template = env.get_template('index.html.j2')
    return template.render(**data)


def get_examples():
    for path in ROOT.iterdir():
        if not path.is_dir():
            continue
        if not (path / 'main.go').exists():
            continue
        if path.name in ('server', 'build', 'frontend'):
            continue
        yield path


def main(args) -> int:
    if args.output:
        build_path = Path(args.output).resolve()
    else:
        build_path = Path(__file__).absolute().parent.parent / 'build'
    build_path.mkdir(exist_ok=True)

    (build_path / 'index.html').write_text(make_index())

    for path in get_examples():
        cmd = [str(path.parent / 'build.sh'), path.name]
        result = subprocess.run(cmd)
        if result.returncode != 0:
            return 1
        src = path.parent / 'build'
        assert src.exists()
        dst = build_path / path.name
        if dst.exists():
            rmtree(str(dst))
        copytree(src=str(src), dst=str(dst))

    print(build_path)
    return 0


if __name__ == '__main__':
    args = parser.parse_args()
    exit(main(args))
