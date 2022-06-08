import mongoengine as me
from .character_model import Character
from .match_model import Match

class Player(me.Document):
  slug = me.StringField(required=True, unique=True)
  username = me.StringField(required=True)
  rating = me.IntField(required=True, default=1200)
  score = me.IntField(required=True, default=0)
  country = me.StringField()
  twitter = me.StringField()
  instagram = me.StringField()
  twitch = me.StringField()
  trovo = me.StringField()
  youtube = me.URLField()
  picture = me.URLField()
  controller = me.StringField()
  real_name = me.StringField()
  main_character = me.ReferenceField(Character)
  sub_characters = me.ListField(me.ReferenceField(Character))
  matches = me.ListField(me.ReferenceField(Match), required=True)