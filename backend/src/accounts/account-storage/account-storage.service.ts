import { Injectable, Scope } from '@nestjs/common';

import { AccountsService } from '../accounts.service';
import { Account } from '../entities/account.entity';

@Injectable({ scope: Scope.REQUEST })
export class AccountStorageService {
  account?: Account;

  constructor(private accountService: AccountsService) {}

  async setByToken(token: string) {
    this.account = await this.accountService.findOne(token);
  }
}
