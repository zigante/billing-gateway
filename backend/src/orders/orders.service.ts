import { Inject, Injectable } from '@nestjs/common';
import { Producer } from '@nestjs/microservices/external/kafka.interface';
import { InjectModel } from '@nestjs/sequelize';
import { EmptyResultError } from 'sequelize';
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
    @Inject('KAFKA_PRODUCER')
    private kafkaProducer: Producer,
  ) {}

  async create(createOrderDto: CreateOrderDto) {
    const order = this.orderModule.create({
      ...createOrderDto,
      accountId: this.accountStorageService.account.id,
    });

    this.kafkaProducer.send({
      topic: 'transactions',
      messages: [{ value: JSON.stringify({ ...createOrderDto, ...order }), key: 'transactions' }],
    });

    return order;
  }

  findAll() {
    return this.orderModule.findAll({ where: { accountId: this.accountStorageService.account.id } });
  }

  findOne(id: string) {
    return this.orderModule.findByPk(id, {
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
