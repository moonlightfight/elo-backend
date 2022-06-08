import email
from flask import Blueprint, jsonify, request
import mongoengine as me
import bcrypt
import jwt
import json
from dotenv import load_dotenv
import os
from ..models.admin_model import Admin

load_dotenv()
secret = os.getenv("SECRET")

admin_routes = Blueprint('admin_routes', __name__)

@admin_routes.route('/create', methods=['POST'])
def create_admin():
  post_data = request.get_json()
  password = post_data.get('password')
  hashed_password = bcrypt.hashpw(password.encode('utf8'), bcrypt.gensalt())
  new_admin = Admin(
    username=post_data.get('username'),
    password=hashed_password,
    email=post_data.get('email')
  )
  new_admin.save()
  return jsonify({
    'status': 'success',
    'message': 'admin created'
  })