import {Module, Global} from '@nestjs/common';
import { ConfigService } from '@nestjs/config';
import * as mysql from 'mysql2/promise';

@Global()
@Module({
  providers: [
    {
      provide: 'MYSQL_CONNECTION',
      inject: [ConfigService],
      useFactory: async (configService: ConfigService) => {
        const pool = mysql.createPool({
          host: configService.get<string>('DB_HOST'), // or the container name if using docker-compose
          port: configService.get<number>('DB_PORT'),
          user: configService.get<string>('DB_USER'),
          password: configService.get<string>('DB_PASSWORD'),
          database: configService.get<string>('DB_NAME'),
          connectionLimit: 10,
        });
        return pool;
      },
    },
  ],
  exports: ['MYSQL_CONNECTION'],
})
export class DatabaseModule {}