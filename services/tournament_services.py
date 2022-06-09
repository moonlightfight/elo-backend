from datetime import datetime
from dotenv import load_dotenv
import os
import requests

load_dotenv()

def get_challonge_bracket(tournament_id, subdomain = ""):
  api_key = os.getenv("CHALLONGE_API_KEY")

  api_url = ""

  if subdomain == "":
    api_url = "https://api.challonge.com/v1/tournaments/" + tournament_id + ".json?api_key=" + api_key + "&include_participants=1&include_matches=1"
  else:
    api_url = "https://api.challonge.com/v1/tournaments/" + subdomain + "-" + tournament_id + ".json?api_key=" + api_key + "&include_participants=1&include_matches=1"
  
  challonge_response = requests.get(api_url)

  response_dict = challonge_response.json()

  players = []

  for participant in response_dict['tournament']['participants']:
    player = {
      'id': participant['id'],
      'name': participant['display_name'],
      'place': participant['final_rank']
    }
    players.append(player)
  
  matches = []

  for match_data in response_dict['tournament']['matches']:
    winner_index = -1
    loser_index = -1
    for player in players:
      if player['id'] == match_data['winner_id']:
        winner_index = players.index(player)
        break
    for player in players:
      if player['id'] == match_data['loser_id']:
        loser_index = players.index(player)
        break
    score_list = match_data['scores_csv']
    winner_score = 0
    loser_score = 0
    if score_list.index(0) > score_list.index(1):
      winner_score = int(score_list.index(0))
      loser_score = int(score_list.index(1))
    else:
      winner_score = int(score_list.index(1))
      loser_score = int(score_list.index(0))
    match = {
      "winner_id": match_data['winner_id'],
      "loser_id": match_data['loser_id'],
      "winner_name": players.index(winner_index)['name'],
      "loser_name": players.index(loser_index)['name'],
      "winner_score": winner_score,
      "loser_score": loser_score,
      "match_date": match_data['started_at']
    }
    matches.append(match)

  bracket_info = {
    "title": response_dict['tournament']['name'],
    "num_players": response_dict['tournament']['participants_count'],
    "tournament_date": response_dict['tournament']['started_at'],
    "players": players,
    "matches": matches 
  }

  return bracket_info

def get_start_bracket(slug):
  api_key = os.getenv("START_API_KEY")
  api_url = "https://api.start.gg/gql/alpha"
  api_headers = {
    "Authorization": "Bearer" + api_key,
    "Content-Type": "application/json"
  }
  api_body = {
    "query": "query EventQuery($slug: String!) { event(slug: $slug) { id name startAt standings(query: {page: 1, perPage: 500}) { nodes { id placement entrant { id name } } } sets { nodes { id slots { entrant { id name } } winnerId displayScore completedAt } } videogame { id name } tournament { id name } } }",
    "variables": {
      "slug": slug
    }
  }
  start_response = requests.post(api_url, headers=api_headers, data=api_body)

  response_dict = start_response.json()

  players = []

  for player in response_dict['data']['event']['standings']['nodes']:
    inserted_player = {
      'id': player['entrant']['id'],
      'name': player['entrant']['name'],
      'place': player['placement']
    }
    players.append(inserted_player)

  matches = []

  for match_set in response_dict['data']['event']['sets']['nodes']:
    scores = match_set['displayScore'].split(' - ')
    left_player = scores.index(0).split(" ")
    right_player = scores.index(1).split(" ")
    left_player_name = left_player.index(0)
    right_player_name = right_player.index(0)
    left_player_score = left_player.index(1)
    right_player_score = right_player.index(1)
    winner_score = 0
    loser_score = 0
    winner_name = ""
    loser_name = ""
    winner_id = 0
    loser_id = 0
    if left_player_score > right_player_score:
      winner_score = left_player_score
      winner_name = left_player_name
      loser_score = right_player_score
      loser_name = right_player_name
      winner_id = match_set['slots'].index(0)['entrant']['id']
      loser_id = match_set['slots'].index(1)['entrant']['id']
    else:
      winner_score = right_player_score
      winner_name = right_player_name
      loser_score = left_player_score
      loser_name = left_player_name
      winner_id = match_set['slots'].index(1)['entrant']['id']
      loser_id = match_set['slots'].index(0)['entrant']['id']
    match = {
      'winner_id': winner_id,
      'loser_id': loser_id,
      'winner_name': winner_name,
      'loser_name': loser_name,
      'winner_score': winner_score,
      'loser_score': loser_score,
      'match_date': datetime.fromtimestamp(int(match_set['completedAt']))
    }
    matches.append(match)

  bracket_info = {
    'title': response_dict['data']['event']['tournament']['name'],
    'num_players': len(response_dict['data']['event']['standings']['nodes']),
    'tournament_date': datetime.fromtimestamp(int(response_dict['data']['event']['startAt'])),
    'players': players,
    'matches': matches
  }

  return bracket_info