import { BelongsTo, Column, DataType, ForeignKey, Model, PrimaryKey, Table } from 'sequelize-typescript';

import { Account } from '../../accounts/entities/account.entity';

export enum OrderStatus {
  PENDING = 'pending',
  APPROVED = 'approved',
  REJECTED = 'rejected',
}

@Table({
  tableName: 'ORDERS',
  createdAt: 'CREATED_AT',
  updatedAt: 'UPDATED_AT',
})
export class Order extends Model {
  @PrimaryKey
  @Column({ type: DataType.UUID, defaultValue: DataType.UUIDV4, field: 'ID' })
  id: string;

  @Column({ type: DataType.DECIMAL(10, 2), allowNull: false, field: 'AMOUNT' })
  amount: number;

  @Column({ allowNull: false, field: 'CREDIT_CARD_NUMBER' })
  creditCardNumber: string;

  @Column({ allowNull: false, field: 'CREDIT_CARD_NAME' })
  creditCardName: string;

  @Column({ allowNull: false, defaultValue: OrderStatus.PENDING, field: 'STATUS' })
  status: OrderStatus;

  @ForeignKey(() => Account)
  @Column({ type: DataType.UUID, allowNull: false, field: 'ACCOUNT_ID' })
  accountId: string;

  @BelongsTo(() => Account)
  account: Account;
}
