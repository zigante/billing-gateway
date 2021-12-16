import { Injectable } from '@nestjs/common';
import { InjectModel } from '@nestjs/sequelize';
import { EmptyResultError, Op } from 'sequelize';
import { AccountStorageService } from 'src/accounts/account-storage/account-storage.service';

import { CreateOrderDto } from './dto/create-order.dto';
import { UpdateOrderDto } from './dto/update-order.dto';
import { Order } from './entities/order.entity';

@Injectable()
export class OrdersService {
  constructor(
    @InjectModel(Order)
    private orderModule: typeof Order,
    private accountStorageService: AccountStorageService,
  ) {}

  create(createOrderDto: CreateOrderDto) {
    return this.orderModule.create({
      ...createOrderDto,
      accountId: this.accountStorageService.account.id,
    });
  }

  findAll() {
    return this.orderModule.findAll({ where: { accountId: this.accountStorageService.account.id } });
  }

  findOne(id: string) {
    return this.orderModule.findOne({
      where: {
        [Op.or]: { id, accountId: this.accountStorageService.account.id },
      },
      rejectOnEmpty: new EmptyResultError(`Account with Id ${id} not found`),
    });
  }

  async update(id: string, updateOrderDto: UpdateOrderDto) {
    const order = await this.findOne(id);
    return order.update(updateOrderDto);
  }

  async remove(id: string) {
    const order = await this.findOne(id);
    return order.destroy();
  }
}
