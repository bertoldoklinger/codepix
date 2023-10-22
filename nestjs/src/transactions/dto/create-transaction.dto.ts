import {
  IsIn,
  IsNotEmpty,
  IsNumber,
  IsOptional,
  IsString,
  Min,
} from 'class-validator';
import { PixKeyKind } from 'src/pix-keys/entities/pix-key.entity';

export class CreateTransactionDto {
  @IsString()
  @IsNotEmpty()
  pix_key_key: string;

  @IsString()
  @IsNotEmpty()
  @IsIn(['cpf', 'email'])
  pix_key_kind: PixKeyKind;

  @IsOptional()
  @IsString()
  description: string = null;

  @IsNumber({ maxDecimalPlaces: 2 })
  @Min(0.01)
  @IsString()
  @IsNotEmpty()
  amount: number;
}
