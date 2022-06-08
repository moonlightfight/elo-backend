import mongoengine as me
from .player_model import Player

class Match(me.Document):
  winning_player = me.ReferenceField(Player, required=True)
  losing_player = me.ReferenceField(Player, required=True)
  date = me.DateTimeField(required=True)
  winner_score = me.IntField(required=True)
  loser_score = me.IntField(required=True)
  is_disqualification = me.BooleanField(required=True, default=False)
  winning_player_starting_elo = me.IntField(required=True)
  winning_player_ending_elo = me.IntField(required=True)
  losing_player_starting_elo = me.IntField(required=True)
  losing_player_ending_elo = me.IntField(required=True)
  