from fastapi import Request
from fastapi.responses import JSONResponse
from fastapi.exceptions import HTTPException
from core.security import verify_token

async def jwt_middleware(request: Request, call_next):
    if request.url.path.startswith("/api/login") or request.url.path.startswith("/public") or request.url.path.startswith("/docs"):
        return await call_next(request)
    auth_header = request.headers.get("Authorization")
    if not auth_header or not auth_header.startswith("Bearer "):
        return JSONResponse(status_code=401, content={"message": "Missing or invalid Authorization header"})

    token = auth_header.split(" ")[1]
    payload = await verify_token(token)
    if not payload:
        return JSONResponse(status_code=401, content={"message": "Invalid or expired token"})
    
    return await call_next(request)
