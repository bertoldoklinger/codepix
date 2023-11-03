import { Box } from "@mui/material";
import { WithdrawForm } from "./WithdrawFrom";




export default async function WithdrawPage({params}: {params: { bankAccountId: string}}) {
  return (
    <Box>
      <WithdrawForm bankAccountId={params.bankAccountId}/>
    </Box>
  )
}