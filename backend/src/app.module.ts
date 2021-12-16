import { Module } from '@nestjs/common';
import { ConfigModule } from '@nestjs/config';
import { SequelizeModule } from '@nestjs/sequelize';

import { AccountsModule } from './accounts/accounts.module';
import { Account } from './accounts/entities/account.entity';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { Order } from './orders/entities/order.entity';
import { OrdersModule } from './orders/orders.module';

@Module({
  imports: [
    ConfigModule.forRoot(),
    OrdersModule,
    AccountsModule,
    SequelizeModule.forRoot({
      autoLoadModels: true,
      database: process.env.DB_DATABASE,
      dialect: process.env.DB_CONNECTION as any,
      host: process.env.DB_HOST,
      models: [Order, Account],
      password: process.env.DB_PASSWORD,
      port: +process.env.DB_PORT,
      username: process.env.DB_USERNAME,
      sync: { alter: true },
    }),
  ],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
