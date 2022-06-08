import mongoengine as me

image_regex = '^[A-Za-z]*\.[a-z]{3}$'

class Character(me.Document):
  name = me.StringField(required=True, unique=True)
  image_link = me.StringField(required=True, unique=True, regex=image_regex)