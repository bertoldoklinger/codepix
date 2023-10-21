import {
  Column,
  CreateDateColumn,
  Entity,
  PrimaryGeneratedColumn,
} from 'typeorm';

//decorator com nome da tabela
@Entity({ name: 'bank-accounts' })
export class BankAccount {
  @PrimaryGeneratedColumn('uuid')
  id: string;

  @Column()
  account_number: string;

  @Column()
  owner_name: string;

  @Column({ default: 0 })
  balance: number;

  @CreateDateColumn({ type: 'timestamp' })
  created_at: Date;
}
