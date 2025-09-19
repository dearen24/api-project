import { Module } from '@nestjs/common';
import { UserController } from './app.controller';
import { UserService } from './app.service';
import { ConfigModule } from '@nestjs/config';
import { DatabaseModule } from './database.module';
import { AuthModule } from './auth/auth.module';

@Module({
  imports: [
    ConfigModule.forRoot({
      isGlobal: true,
      envFilePath: ['../.env'],
    }),
    DatabaseModule,
    AuthModule,
  ],
  controllers: [UserController],
  providers: [UserService],
})
export class AppModule {}
