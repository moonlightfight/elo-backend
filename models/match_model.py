import mongoengine as me
from .player_model import Player

class Match(me.Document):
  winning_player = me.ReferenceField(Player, required=True, db_field="winningPlayer")
  losing_player = me.ReferenceField(Player, required=True, db_field="losingPlayer")
  date = me.DateTimeField(required=True)
  winner_score = me.IntField(required=True, db_field="winnerScore")
  loser_score = me.IntField(required=True, db_field="loserScore")
  is_disqualification = me.BooleanField(required=True, default=False, db_field="isDisqualification")
  winning_player_starting_elo = me.IntField(required=True, db_field="winningPlayerStartingELO")
  winning_player_ending_elo = me.IntField(required=True, db_field="winningPlayerEndingELO")
  losing_player_starting_elo = me.IntField(required=True, db_field="losingPlayerStartingELO")
  losing_player_ending_elo = me.IntField(required=True, db_field="losingPlayerEndingELO")
