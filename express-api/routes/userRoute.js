import express from 'express';
import {getUser, getUsers, getOneUser, createUser, editUser, deleteUser} from '../model/user.js';
import { authenticateToken } from '../middleware/auth.js';
import jwt from 'jsonwebtoken';

const userRouter = express.Router();

userRouter.post('/login', async (req, res) => {
  const { username, password } = req.body;
  const user = await getUser(username);
  if(!user) return res.status(400).json({"message": "User not found"});
  if (password!==user[0].password_hash) return res.status(401).json({ message: "Invalid password" });

  const token = jwt.sign({
      id: user[0].id,
      username: user[0].username
    },
      process.env.JWT_SECRET,
    {
      expiresIn: "1h"
    }
  );

  res.json({token});
});

userRouter.get('/users', authenticateToken, async (req, res) => {
    const users = await getUsers();
    if(!users)res.status(500).json({"message":"Error fetching data"});
    res.json(users);
});

userRouter.get('/users/:id', authenticateToken, async (req, res) => {
    const user = await getOneUser(req.params.id);
    if(!user)res.status(500).json({"message":"Error fetching data"});
    res.json(user);
});

userRouter.post('/users', authenticateToken, async (req, res) => {
  const username = req.body.username;
  const email = req.body.email;
  const password = req.body.password;
  const firstName = req.body.firstName;
  const lastName = req.body.lastName;
  const isActive = req.body.isActive;
  const response = await createUser(username, email, password, firstName, lastName, isActive);
  if(!response)res.status(500).json({"message":"Error adding user"});
  res.status(201).json(response);
});

userRouter.put('/users/:id', authenticateToken, async (req, res) => {
  const username = req.body.username;
  const email = req.body.email;
  const password = req.body.password;
  const firstName = req.body.firstName;
  const lastName = req.body.lastName;
  const isActive = req.body.isActive;
  const response = await editUser(req.params.id, username, email, password, firstName, lastName, isActive);
  if(!response)res.status(500).json({"message":"Error updating user"});
  res.status(200).json(response);
});

userRouter.delete('/users/:id', authenticateToken, async (req, res) => {
  const response = await deleteUser(req.params.id);
  if(!response)res.status(500).json({"message":"Error deleting user"});
  res.status(200).json(response);
});

export default userRouter;