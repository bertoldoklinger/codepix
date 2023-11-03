'use client'

import { Transaction } from "@/models";
import { green, red } from "@mui/material/colors";
import { DataGrid, type GridColDef } from "@mui/x-data-grid";



const columns: GridColDef[] = [
  { field: 'id', headerName: "ID", width: 300},
  { field: 'created_at', headerName: "Data", width: 200, renderCell: (params) => new Date(params.value as string).toLocaleString()},
  { field: 'description', headerName: "Descrição", width: 130},
  { field: 'amount', headerName: "Valor (R$)", width: 180, renderCell: (params) => {
    const amount = new Intl.NumberFormat("pt-BR", {
      style: 'currency',
      currency: 'BRL',
    }).format(params.value as number);
    return (
      <span style={{ color: params.value < 0 ? red[500] : green[500] }}>
        {amount}
      </span>
    );
    }
  },
]

export type MyLatestTransactionsProps = {
  transactions : Transaction[]
  page?: number
  perPage?: number
}

export function MyLatestTransactions({ transactions }: MyLatestTransactionsProps) {
  return <DataGrid rows={transactions} columns={columns}/>
}