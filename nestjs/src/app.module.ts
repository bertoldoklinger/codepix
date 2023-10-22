import { Module } from '@nestjs/common';
import { ConfigModule, ConfigService } from '@nestjs/config';
import { TypeOrmModule } from '@nestjs/typeorm';
import { ConsoleModule } from 'nestjs-console';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { BankAccountsModule } from './bank-accounts/bank-accounts.module';
import { BankAccount } from './bank-accounts/entities/bank-account.entity';
import { FixturesModule } from './fixtures/fixtures.module';
import { PixKey } from './pix-keys/entities/pix-key.entity';
import { PixKeysModule } from './pix-keys/pix-keys.module';
import { Transaction } from './transactions/entities/transaction.entity';
import { TransactionsModule } from './transactions/transactions.module';

@Module({
  imports: [
    ConfigModule.forRoot({
      envFilePath: ['.env', `.bank-${process.env.BANK_CODE}.env`],
      isGlobal: true,
    }),
    ConsoleModule,
    TypeOrmModule.forRootAsync({
      useFactory: (configService: ConfigService) => ({
        type: configService.get('TYPEORM_CONNECTION') as any,
        host: configService.get('TYPEORM_HOST'),
        port: parseInt(configService.get('TYPEORM_PORT')),
        username: configService.get('TYPEORM_USERNAME'),
        password: configService.get('TYPEORM_PASSWORD'),
        database: configService.get('TYPEORM_DATABASE') as string,
        entities: [BankAccount, PixKey, Transaction],
        synchronize: true,
      }),
      inject: [ConfigService],
    }),
    BankAccountsModule,
    PixKeysModule,
    TransactionsModule,
    FixturesModule,
  ],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
