from models.player_model import Player
from flask import Blueprint, jsonify, request

from services.utils import generate_slug

player_routes = Blueprint('player_routes', __name__)


@player_routes.route('/create', methods=["POST"])
def create_player():
    get_data = request.get_json()
    player = get_data.get('player')
    slug = generate_slug(player.name)
