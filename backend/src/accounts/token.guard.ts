import { CanActivate, ExecutionContext, Injectable } from '@nestjs/common';

import { AccountStorageService } from './account-storage/account-storage.service';

@Injectable()
export class TokenGuard implements CanActivate {
  constructor(private accountStorageService: AccountStorageService) {}

  async canActivate(context: ExecutionContext) {
    if (context.getType() !== 'http') return true;

    const request = context.switchToHttp().getRequest();
    const token = request.headers?.['x-api-token'];

    if (token) {
      try {
        await this.accountStorageService.setByToken(token);
        return true;
      } catch (e) {
        console.error(e);
        return false;
      }
    }

    return false;
  }
}
