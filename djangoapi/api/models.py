# This is an auto-generated Django model module.
# You'll have to do the following manually to clean this up:
#   * Rearrange models' order
#   * Make sure each model has one field with primary_key=True
#   * Make sure each ForeignKey and OneToOneField has `on_delete` set to the desired behavior
#   * Remove `managed = False` lines if you wish to allow Django to create, modify, and delete the table
# Feel free to rename the models, but don't rename db_table values or field names.
from django.db import models
from django.contrib.auth.models import AbstractBaseUser, BaseUserManager, PermissionsMixin

class ApiKeys(models.Model):
    user = models.ForeignKey('Users', models.DO_NOTHING)
    key_hash = models.CharField(max_length=255)
    name = models.CharField(max_length=100)
    permissions = models.JSONField()
    rate_limit = models.IntegerField(blank=True, null=True)
    is_active = models.IntegerField(blank=True, null=True)
    expires_at = models.DateTimeField(blank=True, null=True)
    last_used_at = models.DateTimeField(blank=True, null=True)
    created_at = models.DateTimeField(blank=True, null=True)

    class Meta:
        managed = False
        db_table = 'api_keys'


class Categories(models.Model):
    name = models.CharField(unique=True, max_length=100)
    description = models.TextField(blank=True, null=True)
    color = models.CharField(max_length=7, blank=True, null=True)
    created_at = models.DateTimeField(blank=True, null=True)

    class Meta:
        managed = False
        db_table = 'categories'


class Comments(models.Model):
    post = models.ForeignKey('Posts', models.DO_NOTHING)
    user = models.ForeignKey('Users', models.DO_NOTHING)
    content = models.TextField()
    parent_comment = models.ForeignKey('self', models.DO_NOTHING, blank=True, null=True)
    is_approved = models.IntegerField(blank=True, null=True)
    created_at = models.DateTimeField(blank=True, null=True)
    updated_at = models.DateTimeField(blank=True, null=True)

    class Meta:
        managed = False
        db_table = 'comments'


class PostTags(models.Model):
    pk = models.CompositePrimaryKey('post_id', 'tag_id')
    post = models.ForeignKey('Posts', models.DO_NOTHING)
    tag = models.ForeignKey('Tags', models.DO_NOTHING)
    created_at = models.DateTimeField(blank=True, null=True)

    class Meta:
        managed = False
        db_table = 'post_tags'


class Posts(models.Model):
    title = models.CharField(max_length=255)
    content = models.TextField()
    excerpt = models.TextField(blank=True, null=True)
    author = models.ForeignKey('Users', models.DO_NOTHING)
    category = models.ForeignKey(Categories, models.DO_NOTHING, blank=True, null=True)
    status = models.CharField(max_length=9, blank=True, null=True)
    view_count = models.IntegerField(blank=True, null=True)
    is_featured = models.IntegerField(blank=True, null=True)
    published_at = models.DateTimeField(blank=True, null=True)
    created_at = models.DateTimeField(blank=True, null=True)
    updated_at = models.DateTimeField(blank=True, null=True)

    class Meta:
        managed = False
        db_table = 'posts'


class Tags(models.Model):
    name = models.CharField(unique=True, max_length=100)
    slug = models.CharField(unique=True, max_length=100)
    created_at = models.DateTimeField(blank=True, null=True)

    class Meta:
        managed = False
        db_table = 'tags'


class UserManager(BaseUserManager):
    def create_user(self, username, email, first_name, last_name, is_active, password_hash=None, **extra_fields):
        if not email:
            raise ValueError("Users must have an email address")
        if not username:
            raise ValueError("Users must have a username")

        email = self.normalize_email(email)
        user = self.model(username=username, email=email, first_name=first_name, last_name=last_name, is_active=is_active, **extra_fields)
        user.set_password(password_hash)  # hashes the password
        user.save(using=self._db)
        return user

    

class Users(AbstractBaseUser):
    id = models.AutoField(primary_key=True)
    username = models.CharField(unique=True, max_length=50)
    email = models.CharField(unique=True, max_length=100)
    password_hash = models.CharField(max_length=255)
    first_name = models.CharField(max_length=50, blank=True, null=True)
    last_name = models.CharField(max_length=50, blank=True, null=True)
    is_active = models.IntegerField(blank=True, null=True)
    created_at = models.DateTimeField(auto_now_add=True)
    updated_at = models.DateTimeField(auto_now=True)

    @property
    def password(self):
        return self.password_hash

    def set_password(self, raw_password):
        from django.contrib.auth.hashers import make_password
        self.password_hash = make_password(raw_password)

    objects = UserManager()

    USERNAME_FIELD = "username"   # field used for login
    REQUIRED_FIELDS = ["email"]   # required when creating superuser

    class Meta:
        db_table = "users"  # IMPORTANT: match your MySQL table name

    def __str__(self):
        return self.username