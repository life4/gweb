import yaml
from jinja2 import Environment, FileSystemLoader
from pathlib import Path


env = Environment(
    loader=FileSystemLoader('.'),
    extensions=['jinja2_markdown.MarkdownExtension'],
)


def make_index():
    data = yaml.safe_load(Path('index.yml').open())
    template = env.get_template('index.html.j2')
    return template.render(**data)


if __name__ == '__main__':
    build_path = Path(__file__).absolute().parent.parent / 'build'
    (build_path / 'index.html').write_text(make_index())
    print(build_path)
