import { Module } from '@nestjs/common';
import { SequelizeModule } from '@nestjs/sequelize';

import { AccountsModule } from '../accounts/accounts.module';
import { Order } from './entities/order.entity';
import { OrdersController } from './orders.controller';
import { OrdersService } from './orders.service';

@Module({
  imports: [SequelizeModule.forFeature([Order]), AccountsModule],
  controllers: [OrdersController],
  providers: [OrdersService],
})
export class OrdersModule {}
