from fastapi import FastAPI, Depends, HTTPException, status, Header
from fastapi.security import HTTPBasic, HTTPBasicCredentials
from fastapi.responses import JSONResponse
from sqlalchemy.orm import Session
from typing import List, Optional
from database import engine, get_db
from models import Base, User

async def get_users_db(db, skip: int = 0, limit: int = 10):
    users = db.query(User).offset(skip).limit(limit).all()
    return users

async def get_user_db(db, user_id: int):
    user = db.query(User).filter(User.id == user_id).first()
    return user

async def create_user_db(db, user_data: dict):
    user = User(**user_data.dict())
    db.add(user)
    try:
        db.commit()
        db.refresh(user)
        return user
    except Exception as e:
        db.rollback()
        return JSONResponse(status_code=400, content={"message": f"Error creating user: {str(e)}"})

async def update_user_db(db, user_id: int, user_data: dict):
    user = db.query(User).filter(User.id == user_id).first()
    update_data = user_data.model_dump(exclude_unset=True)
    try:
        for key, value in update_data.items():
            setattr(user, key, value)
        db.commit()
        db.refresh(user)
        return user
    except Exception as e:
        db.rollback()
        return JSONResponse(status_code=400, content={"message": f"Error updating user: {str(e)}"})

async def delete_user_db(db, user_id: int):
    user = db.query(User).filter(User.id == user_id).first()
    if user is None:
        return JSONResponse(status_code=404, content={"message": "User not found"})
    try:
        db.delete(user)
        db.commit()
        return JSONResponse(status_code=200, content={"message": "User deleted successfully"})
    except Exception as e:
        db.rollback()
        return JSONResponse(status_code=400, content={"message": f"Error deleting user: {str(e)}"})