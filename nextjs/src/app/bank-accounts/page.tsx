import { CardAction } from "@/components/CardAction";
import { BankAccount } from "@/models";
import { Typography } from "@mui/material";
import Grid2 from "@mui/material/Unstable_Grid2/Grid2";

export async function getBankAccounts(): Promise<BankAccount[]> {
  const response = await fetch('http://host.docker.internal:3000/bank-accounts')
  return response.json()
}

export async function HomePage() {
  const bankAccounts = await getBankAccounts()
  return (
    <div>
      <Typography variant="h4">Contas bancárias</Typography>

      <Grid2 container spacing={2} mt={1}>
        {bankAccounts.map((bankAccount)=> 
          <Grid2 key={bankAccount.id} xs={12} sm={6} md={4}>
            <CardAction >
              <Typography variant="h5" component="div">
              {bankAccount.owner_name} 
              </Typography>
              <Typography component='span'>
              Conta: {bankAccount.account_number}
              </Typography>
            </CardAction>
          </Grid2>
        )}
      </Grid2>
    </div>
  )
}

export default HomePage;