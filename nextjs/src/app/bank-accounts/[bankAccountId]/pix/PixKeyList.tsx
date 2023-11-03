import { PixKey } from "@/models";
import { StarBorder } from "@mui/icons-material";
import { List, ListItem, ListItemButton, ListItemIcon, ListItemText, Typography } from "@mui/material";



export async function getPixKeys(bankAccountId:string):Promise<PixKey[]> {
  const response = await fetch(`${process.env.NEXT_PUBLIC_NEST_API_URL}/bank-accounts/${bankAccountId}/pix-keys`,
    { 
      next: {
        tags: [`pix-keys-${bankAccountId}`]
      }
    }
  )

  return response.json()
}

export type PixKeyListProps = {
  bankAccountId: string;
}

export async function PixKeyList({ bankAccountId }: PixKeyListProps) {
  const pixKeys = await getPixKeys(bankAccountId)
  return (
    <div>
      <Typography variant="h5">Minhas chaves pix</Typography>
      <List>
        {
          pixKeys.map((pixKey)=> (
            <ListItem disablePadding key={pixKey.id}>
              <ListItemButton>
                <ListItemIcon>
                  <StarBorder />
                </ListItemIcon>
                <ListItemText primary={pixKey.key} secondary={pixKey.kind}/>
              </ListItemButton>
            </ListItem>
          ))
        }
      </List>
    </div>
  )
}