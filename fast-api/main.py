from fastapi import FastAPI, Depends, HTTPException, status, Header
from fastapi.security import HTTPBasic, HTTPBasicCredentials
import secrets
from sqlalchemy.orm import Session
from typing import List, Optional
from pydantic import BaseModel
from database import engine, get_db
from models import Base, User, Post

# # Print what tables SQLAlchemy knows about
# print("🔍 SQLAlchemy knows about these tables:")
# for table_name, table in Base.metadata.tables.items():
#     print(f"  - {table_name}")
    
# if not Base.metadata.tables:
#     print("❌ No tables found! Models might not be imported correctly.")

# # Create tables with error handling
# try:
#     print("🔄 Creating database tables...")
#     Base.metadata.create_all(bind=engine)
#     print("✅ Tables created successfully!")
# except Exception as e:
#     print(f"❌ Error creating tables: {e}")
#     print("Make sure your database exists and credentials are correct")

app = FastAPI(title="FastAPI MySQL Example")
security =  HTTPBasic()

# Pydantic schemas
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
    is_active: Optional[bool] = None
    password_hash: Optional[str] = None
    first_name: Optional[str] = None
    last_name: Optional[str] = None

class UserResponse(BaseModel):
    id: int
    username: str
    email: str
    first_name: str
    last_name: str
    is_active: bool
    
    class Config:
        from_attributes = True

class PostCreate(BaseModel):
    title: str
    content: str = None
    author_id: int

class PostResponse(BaseModel):
    id: int
    title: str
    content: str = None
    author_id: int
    
    class Config:
        from_attributes = True

async def validate_users(username: str, password: str, db: Session = Depends(get_db)):
    user = db.query(User).filter(User.username == username).first()
    if user:
        correct_password = secrets.compare_digest(password, user.password_hash)
        if correct_password:
            return True
        else:
            raise HTTPException(
                status_code=status.HTTP_401_UNAUTHORIZED,
                detail="Incorrect password",
                headers={"WWW-Authenticate": "Basic"},
            )
    else:
        raise HTTPException(
            status_code=status.HTTP_401_UNAUTHORIZED,
            detail="User not found",
            headers={"WWW-Authenticate": "Basic"},
        )

# Health check endpoint
@app.get("/health")
async def health_check():
    return {"status": "healthy", "database": "connected"}

# User endpoints
@app.post("/users/", response_model=UserResponse)
async def create_user(user: UserCreate, db: Session = Depends(get_db), x_username: Optional[str] = Header(alias="x-username"), x_password: Optional[str] = Header(alias="x-password")):
    if(await validate_users(x_username, x_password, db)):
        db_user = User(**user.model_dump())
        db.add(db_user)
        try:
            db.commit()
            db.refresh(db_user)
            return db_user
        except Exception as e:
            db.rollback()
            raise HTTPException(status_code=400, detail=f"Error creating user: {str(e)}")

@app.get("/users/", response_model=List[UserResponse])
async def get_users(skip: int = 0, limit: int = 100, db: Session = Depends(get_db), x_username: Optional[str] = Header(alias="x-username"), x_password: Optional[str] = Header(alias="x-password")):
    if(await validate_users(x_username, x_password, db)):
        users = db.query(User).offset(skip).limit(limit).all()
        return users

@app.get("/users/{user_id}", response_model=UserResponse)
async def get_user(user_id: int, db: Session = Depends(get_db), x_username: Optional[str] = Header(alias="x-username"), x_password: Optional[str] = Header(alias="x-password")):
    if(await validate_users(x_username, x_password, db)):
        user = db.query(User).filter(User.id == user_id).first()
        return user
    
@app.put("/users/{user_id}", response_model=UserResponse)
async def update_user(user_id: int, user: UserUpdate, db: Session = Depends(get_db), x_username: Optional[str] = Header(alias="x-username"), x_password: Optional[str] = Header(alias="x-password")):
    if(await validate_users(x_username, x_password, db)):
        db_user = db.query(User).filter(User.id == user_id).first()
        update_data = user.model_dump(exclude_unset=True)
        for field, value in update_data.items():
            setattr(db_user, field, value)
        try:
            db.commit()
            db.refresh(db_user)
            return db_user
        except Exception as e:
            db.rollback()
            raise HTTPException(status_code=400, detail=f"Error editing user: {str(e)}")

# Post endpoints
@app.post("/posts/", response_model=PostResponse)
async def create_post(post: PostCreate, db: Session = Depends(get_db), x_username: Optional[str] = Header(alias="x-username"), x_password: Optional[str] = Header(alias="x-password")):
    if(await validate_users(x_username, x_password, db)):
        db_post = Post(**post.model_dump())
        db.add(db_post)
        try:
            db.commit()
            db.refresh(db_post)
            return db_post
        except Exception as e:
            db.rollback()
            raise HTTPException(status_code=400, detail=f"Error creating post: {str(e)}")

@app.get("/posts/", response_model=List[PostResponse])
async def get_posts(skip: int = 0, limit: int = 100, db: Session = Depends(get_db), x_username: Optional[str] = Header(alias="x-username"), x_password: Optional[str] = Header(alias="x-password")):
    if(await validate_users(x_username, x_password, db)):
        posts = db.query(Post).offset(skip).limit(limit).all()
        return posts

if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=8000)