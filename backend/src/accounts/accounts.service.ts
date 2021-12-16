import { Injectable } from '@nestjs/common';
import { InjectModel } from '@nestjs/sequelize';

import { CreateAccountDto } from './dto/create-account.dto';
import { UpdateAccountDto } from './dto/update-account.dto';
import { Account } from './entities/account.entity';

@Injectable()
export class AccountsService {
  constructor(@InjectModel(Account) private accountModule: typeof Account) {}

  create(createAccountDto: CreateAccountDto) {
    return this.accountModule.create(createAccountDto);
  }

  findAll() {
    return this.accountModule.findAll();
  }

  findOne(id: string) {
    return this.accountModule.findByPk(id);
  }

  async update(id: string, updateAccountDto: UpdateAccountDto) {
    const account = await this.findOne(id);
    return account.update(updateAccountDto);
  }

  async remove(id: string) {
    const account = await this.findOne(id);
    return account.destroy();
  }
}
