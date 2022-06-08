import mongoengine as me


class Admin(me.Document):
    username = me.StringField(required=True, unique=True)
    password = me.StringField(required=True)
    email = me.EmailField(required=True, unique=True)
