from fastapi import APIRouter, Request
from fastapi import FastAPI, Depends, HTTPException, status, Header
from fastapi.security import HTTPBasic, HTTPBasicCredentials
from pydantic import BaseModel
from sqlalchemy.orm import Session
from typing import List, Optional
from database import engine, get_db
from model.user import get_users_db, get_user_db, update_user_db, create_user_db, delete_user_db
from core.security import create_access_token

router = APIRouter()

class UserCreate(BaseModel):
    username: str
    email: str
    password_hash: str
    first_name: str = None
    last_name: str = None
    is_active: bool = True

class UserUpdate(BaseModel):
    username: Optional[str] = None
    email: Optional[str] = None
    password_hash: Optional[str] = None
    first_name: Optional[str] = None
    last_name: Optional[str] = None
    is_active: Optional[bool] = None

class UserResponse(BaseModel):
    id: int
    username: str
    email: str
    first_name: str
    last_name: str
    is_active: bool
    
    class Config:
        from_attributes = True

# User endpoints
@router.post("/login")
async def login(request: Request):
    body = await request.json()
    token = await create_access_token(body)
    return {"token": token}

@router.post("/users/", response_model=UserResponse)
async def create_user(user: UserCreate, db: Session = Depends(get_db)):
    user = await create_user_db(db, user)
    return user

@router.get("/users", response_model=List[UserResponse])
async def get_users(db: Session = Depends(get_db)):
    users = await get_users_db(db)
    return users

@router.get("/users/{user_id}", response_model=UserResponse)
async def get_user(user_id: int, db: Session = Depends(get_db)):
    user = await get_user_db(db, user_id) 
    if user is None:
        return JSONResponse(status_code=404, content={"message": "User not found"})
    return user
    
@router.put("/users/{user_id}", response_model=UserResponse)
async def update_user(user_id: int, user: UserUpdate, db: Session = Depends(get_db)):
    user = await update_user_db(db, user_id, user)
    return user

@router.delete("/users/{user_id}")
async def delete_user(user_id: int, db: Session = Depends(get_db)):
    result = await delete_user_db(db, user_id)
    return result