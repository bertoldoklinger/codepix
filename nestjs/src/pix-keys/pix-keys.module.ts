import { Module } from '@nestjs/common';
import { ClientsModule, Transport } from '@nestjs/microservices';
import { TypeOrmModule } from '@nestjs/typeorm';
import { join } from 'path';
import { BankAccount } from 'src/bank-accounts/entities/bank-account.entity';
import { PixKey } from './entities/pix-key.entity';
import { PixKeysController } from './pix-keys.controller';
import { PixKeysService } from './pix-keys.service';

@Module({
  imports: [
    TypeOrmModule.forFeature([PixKey, BankAccount]),
    ClientsModule.register([
      {
        name: 'PIX_PACKAGE',
        transport: Transport.GRPC,
        options: {
          url: 'host.docker.internal:50051',
          package: 'github.com.bertoldoklinger.codepix',
          protoPath: join(__dirname, 'proto', 'pix.proto'),
        },
      },
    ]),
  ],
  controllers: [PixKeysController],
  providers: [PixKeysService],
})
export class PixKeysModule {}
