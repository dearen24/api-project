import { Body, Controller, Delete, Get, NotFoundException, Param, Post, Put, UseGuards } from '@nestjs/common';
import { UserService } from './app.service';
import { OkPacket, QueryResult } from 'mysql2';
import { AuthGuard } from '@nestjs/passport';

@Controller('api')
export class UserController {
  constructor(private readonly userService: UserService) {}

  @Get('users')
  @UseGuards(AuthGuard('jwt'))
  async getUsers(): Promise<QueryResult> {
    return await this.userService.findUsers();
  }

  @Get('users/:id')
  @UseGuards(AuthGuard('jwt'))
  async getUser(@Param('id') id: string): Promise<QueryResult> {
    const user = await this.userService.findUserById(Number(id));
    if (!user) {
      throw new NotFoundException(`User with ID ${id} not found`);
    }
    return user;
  }

  @Post('users')
  @UseGuards(AuthGuard('jwt'))
  async addUser(@Body() body: { username: string; email: string; first_name: string; last_name: string; password_hash: string; is_active: number }): Promise<QueryResult> {
    const user = await this.userService.createUser(body.username, body.email, body.first_name, body.last_name, body.password_hash, body.is_active);
    return user;
  }

  @Put('users/:id')
  @UseGuards(AuthGuard('jwt'))
  async updateUser(@Param('id') id: number ,@Body() body: { username: string; email: string; first_name: string; last_name: string; password_hash: string; is_active: number }): Promise<QueryResult> {
    const user = await this.userService.editUser(body.username, body.email, body.first_name, body.last_name, body.password_hash, body.is_active, id);
    return user;
  }

  @Delete('users/:id')
  @UseGuards(AuthGuard('jwt'))
  async removeUser(@Param('id') id: number): Promise<QueryResult> {
    const user = await this.userService.deleteUser(id);
    return user;
  }
}
