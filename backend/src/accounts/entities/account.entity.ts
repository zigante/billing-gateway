import { Column, DataType, Model, PrimaryKey, Table } from 'sequelize-typescript';

@Table({
  tableName: 'ACCOUNTS',
  createdAt: 'CREATED_AT',
  updatedAt: 'UPDATED_AT',
})
export class Account extends Model {
  @PrimaryKey
  @Column({ type: DataType.UUIDV4, defaultValue: DataType.UUIDV4, field: 'ID' })
  id: string;

  @Column({ allowNull: false, field: 'NAME' })
  name: number;

  @Column({ allowNull: false, defaultValue: () => Math.random().toString(36).slice(2), field: 'TOKEN' })
  token: string;
}
