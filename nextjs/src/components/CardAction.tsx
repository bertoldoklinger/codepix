
import Grid2 from "@mui/material/Unstable_Grid2/Grid2";
import { PropsWithChildren } from "react";
import { Card } from "./Card";

export type CardActionProps = {

}



export function CardAction({children, }: PropsWithChildren<CardActionProps>){
  return (
  <Card>
    <Grid2 container spacing={2}>
      <Grid2 xs={12} sm={9}>
          {children}
      </Grid2>
      <Grid2 
      xs={12} 
      sm={3} 
      display={'flex'} 
      alignItems={'center'} 
      justifyContent={'flex-end'}
      >
        Ação
      </Grid2>
    </Grid2>
  </Card>
  )
}