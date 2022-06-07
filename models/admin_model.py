from typing_extensions import Required
import mongoengine as me

class Admin(me.Document):
  username = me.StringField(required=True, unique=True)
  password = me.StringField(required=True)
  email = me.StringField(required=True, unique=True)