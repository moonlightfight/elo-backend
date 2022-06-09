from datetime import datetime
import mongoengine as me
from .player_model import Player
from .character_model import Character
from .match_model import Match

BRACKET_TYPE = ("Double Elim", "Single Elim", "Swiss", "Round Robin", "Hybrid")


class TournamentResult(me.EmbeddedDocument):
    place = me.IntField(required=True)
    points = me.IntField(required=True)
    player = me.ReferenceField(Player, required=True)
    charactersUsed = me.ListField(me.ReferenceField(Character))


class Tournament(me.Document):
    name = me.StringField(required=True, unique=True)
    slug = me.StringField(required=True, unique=True)
    location = me.StringField()
    bracket_url = me.URLField(db_field="bracketUrl")
    no_bracket = me.BooleanField(
        required=True, default=False, db_field="noBracket")
    num_players = me.IntField(required=True, db_field="numPlayers")
    date = me.DateTimeField(required=True)
    date_added = me.DateTimeField(
        required=True, default=datetime.now, db_field="dateAdded")
    replay = me.URLField()
    results = me.ListField(me.EmbeddedDcoumentField(
        TournamentResult), required=True)
    matches = me.ListField(me.ReferenceField(Match), required=True)
    bracket_type = me.EnumField(BRACKET_TYPE, required=True)
