import re


def generate_slug(formatted_name):
    lower_name = formatted_name.lower()
    special_char_regex = re.compile('([^A-Za-z0-9\s_-])')
    spaces_to_dashes = lower_name.replace(' ', '-')
    underscores_to_dashes = spaces_to_dashes.replace('_', '-')
    slug = special_char_regex.sub(underscores_to_dashes, '')
    return slug
