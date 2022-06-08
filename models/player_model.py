import mongoengine as me

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
  realName = me.StringField()