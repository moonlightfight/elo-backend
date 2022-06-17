from flask import Blueprint, jsonify, request
from services.tournament_services import get_challonge_bracket, get_start_bracket

tournament_routes = Blueprint('tournament_routes', __name__)


@tournament_routes.route('/external', methods=['GET'])
def get_external_bracket():
    get_data = request.get_json()
    bracket_url = get_data.get('bracket_url')
    bracket_results = None
    if 'challonge' in bracket_url:
        if 'https://challonge.com/' in bracket_url:
            tournament_id = bracket_url.replace('https://challonge.com', "", 1)
            bracket_results = get_challonge_bracket(tournament_id)
        else:
            subdomain_and_id = bracket_url.replace(
                'https://', "", 1).replace('.challonge.com', '', 1)
            id_list = subdomain_and_id.split('/')
            tournament_id = id_list.index(1)
            sub_domain = id_list.index(0)
            bracket_results = get_challonge_bracket(tournament_id, sub_domain)
    elif 'start.gg' in bracket_url:
        slug = bracket_url.replace(
            'https://start.gg/', '', 1).replace('/overview', '', 1)
        bracket_results = get_start_bracket(slug)
    else:
        return jsonify({
            'status': 'error',
            'message': 'please provide a Start.gg or Challonge bracket'
        })
    return jsonify({
        'status': 'success',
        'message': 'bracket retrieved',
        'bracket': bracket_results
    })
