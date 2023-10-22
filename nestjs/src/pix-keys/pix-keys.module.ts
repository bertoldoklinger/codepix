import { Module } from '@nestjs/common';
import { ConfigService } from '@nestjs/config';
import { ClientsModule, Transport } from '@nestjs/microservices';
import { TypeOrmModule } from '@nestjs/typeorm';
import { join } from 'path';
import { BankAccount } from '../bank-accounts/entities/bank-account.entity';
import { PixKey } from './entities/pix-key.entity';
import { PixKeysController } from './pix-keys.controller';
import { PixKeysService } from './pix-keys.service';

@Module({
  imports: [
    TypeOrmModule.forFeature([PixKey, BankAccount]),
    ClientsModule.registerAsync([
      {
        name: 'PIX_PACKAGE',
        useFactory: (configService: ConfigService) => ({
          transport: Transport.GRPC,
          options: {
            url: configService.get('GRPC_URL'),
            package: 'github.com.bertoldoklinger.codepix',
            protoPath: [join(__dirname, 'proto/pix.proto')],
          },
        }),
        inject: [ConfigService],
      },
    ]),
  ],
  controllers: [PixKeysController],
  providers: [PixKeysService],
})
export class PixKeysModule {}
