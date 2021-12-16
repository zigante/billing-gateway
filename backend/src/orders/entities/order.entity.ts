import { Column, DataType, Model, PrimaryKey, Table } from 'sequelize-typescript';

export enum OrderStatus {
  PENDING = 'pending',
  APPROVED = 'approved',
}

@Table({
  tableName: 'ORDERS',
  createdAt: 'CREATED_AT',
  updatedAt: 'UPDATED_AT',
})
export class Order extends Model {
  @PrimaryKey
  @Column({ type: DataType.UUIDV4, defaultValue: DataType.UUIDV4, field: 'ID' })
  id: string;

  @Column({ type: DataType.DECIMAL, allowNull: false, field: 'AMOUNT' })
  amount: number;

  @Column({ allowNull: false, field: 'CREDIT_CARD_NUMBER' })
  creditCardNumber: string;

  @Column({ allowNull: false, field: 'CREDIT_CARD_NAME' })
  creditCardName: string;

  @Column({ allowNull: false, defaultValue: OrderStatus.PENDING, field: 'STATUS' })
  status: OrderStatus;
}
