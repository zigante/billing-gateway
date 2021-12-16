import { Module } from '@nestjs/common';
import { SequelizeModule } from '@nestjs/sequelize';

import { AccountStorageService } from './account-storage/account-storage.service';
import { AccountsController } from './accounts.controller';
import { AccountsService } from './accounts.service';
import { Account } from './entities/account.entity';
import { TokenGuard } from './token.guard';

@Module({
  imports: [SequelizeModule.forFeature([Account])],
  controllers: [AccountsController],
  providers: [AccountsService, AccountStorageService, TokenGuard],
  exports: [AccountStorageService],
})
export class AccountsModule {}
