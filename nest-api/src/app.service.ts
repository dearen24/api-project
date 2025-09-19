import { Inject, Injectable } from '@nestjs/common';
import * as promise from 'mysql2/promise';

@Injectable()
export class UserService {
    constructor(@Inject('MYSQL_CONNECTION') private pool: promise.Pool) {}

    async findUsers(): Promise<promise.QueryResult> {
      const [rows] = await this.pool.query('SELECT * FROM users');
      return rows;
    }

    async findUserById(id: number): Promise<promise.QueryResult> {
      const [rows] = await this.pool.query(`SELECT * FROM users WHERE id = ${id}`);
      return rows[0];
    }

    async createUser(username: string, email: string, first_name: string, last_name: string, password_hash: string, is_active: number): Promise<promise.QueryResult> {
      const [result] = await this.pool.execute(`INSERT INTO users (username, email, password_hash, first_name, last_name, is_active, created_at, updated_at) VALUES('${username}','${email}','${password_hash}','${first_name}','${last_name}',${is_active}, NOW(), NOW())`);
      return result;
    }

    async editUser(username: string, email: string, first_name: string, last_name: string, password_hash: string, is_active: number, id: number): Promise<promise.QueryResult> {
      const [result] = await this.pool.execute(`UPDATE users SET username = '${username}', email = '${email}', password_hash = '${password_hash}', first_name = '${first_name}', last_name = '${last_name}', is_active = ${is_active}, updated_at = NOW() WHERE id=${id}`);
      return result;
    }

    async deleteUser(id: number): Promise<promise.QueryResult> {
      const [rows] = await this.pool.query(`DELETE FROM users WHERE id=${id}`);
      return rows;
    }
}
