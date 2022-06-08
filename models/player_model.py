from models.tournament_model import Tournament
import mongoengine as me
from .character_model import Character
from .match_model import Match
from .tournament_model import Tournament

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
  real_name = me.StringField(db_field="realName")
  main_character = me.ReferenceField(Character, db_field="mainCharacter")
  sub_characters = me.ListField(me.ReferenceField(Character), db_field="subCharacters")
  matches = me.ListField(me.ReferenceField(Match), required=True)
  tournaments = me.ListField(me.ReferenceField(Tournament), require=True)