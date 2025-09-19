import { Injectable, UnauthorizedException } from '@nestjs/common';
import { JwtService } from '@nestjs/jwt';
import * as bcrypt from 'bcrypt';
import * as promise from 'mysql2/promise';
import { Inject } from '@nestjs/common';

@Injectable()
export class AuthService {
  constructor(
    private jwtService: JwtService,
    @Inject('MYSQL_CONNECTION') private pool: promise.Pool,
  ) {}

  async validateUser(username: string, pass: string) {
    const [rows]: any[] = await this.pool.query(
      'SELECT * FROM users WHERE username = ?',
      [username],
    );
    const user = rows[0];
    if (user && (await bcrypt.compare(pass, user.password_hash))) {
      const { password, ...result } = user;
      return await result;
    }
    return null;
  }

  async login(user: any) {
    const payload = { username: user.username, sub: user.id };
    return {
      access_token: this.jwtService.sign(payload),
    };
  }
}