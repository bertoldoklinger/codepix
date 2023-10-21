import { Injectable } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { Repository } from 'typeorm';
import { CreateBankAccountDto } from './dto/create-bank-account.dto';
import { BankAccount } from './entities/bank-account.entity';

@Injectable()
export class BankAccountsService {
  constructor(
    @InjectRepository(BankAccount)
    private bankAccountRepository: Repository<BankAccount>,
  ) {}

  create(createBankAccountDto: CreateBankAccountDto) {
    return this.bankAccountRepository.save(createBankAccountDto);
  }

  findAll() {
    return this.bankAccountRepository.find();
  }

  findOne(id: string) {
    return this.bankAccountRepository.findOneOrFail({ where: { id } });
  }

  // update(id: number, updateBankAccountDto: UpdateBankAccountDto) {
  //   return `This action updates a #${id} bankAccount`;
  // }

  // remove(id: number) {
  //   return `This action removes a #${id} bankAccount`;
  // }
}
