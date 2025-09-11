from fastapi import FastAPI, Depends, HTTPException, status, Header
from fastapi.security import HTTPBasic, HTTPBasicCredentials
from routes import user
from middleware import auth
import secrets
from sqlalchemy.orm import Session
from typing import List, Optional
from database import engine, get_db
from models import Base, User, Post
from dotenv import load_dotenv
import os
load_dotenv()

app = FastAPI(title="FastAPI MySQL Example")

# Middleware
app.middleware("http")(auth.jwt_middleware)

security =  HTTPBasic()

# Include router
app.include_router(user.router, prefix="/api")

# Health check endpoint
@app.get("/health")
async def health_check():
    return {"status": "healthy", "database": "connected"}

if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=3000)