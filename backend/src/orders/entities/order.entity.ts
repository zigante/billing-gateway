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
  @Column({ type: DataType.UUIDV4, defaultValue: DataType.UUIDV4 })
  id: string;

  @Column({ type: DataType.DECIMAL, allowNull: false })
  amount: number;

  @Column({ allowNull: false })
  creditCardNumber: string;

  @Column({ allowNull: false })
  creditCardName: string;

  @Column({ allowNull: false, defaultValue: OrderStatus.PENDING })
  status: OrderStatus;
}
