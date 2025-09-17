from rest_framework import serializers
from .models import Users

class UserSerializer(serializers.ModelSerializer):
    class Meta:
        model = Users
        fields = '__all__'

    def create(self, validated_data):
        user = Users.objects.create_user(
            username=validated_data["username"],
            email=validated_data["email"],
            password_hash=validated_data["password_hash"],
            first_name=validated_data["first_name"],
            last_name=validated_data["last_name"],
            is_active=validated_data["is_active"],
        )
        return user

    def update(self, instance, validated_data):
        password = validated_data.pop("password_hash", None)

        # update normal fields
        for attr, value in validated_data.items():
            setattr(instance, attr, value)

        # handle password separately
        if password:
            instance.set_password(password)

        instance.save()
        return instance