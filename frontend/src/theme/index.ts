import { createTheme, Theme } from '@mui/material/styles';

declare module '@mui/material/styles' {
  interface Theme {
    custom: {
      sidebar: {
        width: number;
        background: string;
      };
      header: {
        height: number;
        background: string;
      };
    }
  }

  interface ThemeOptions {
    custom?: {
      sidebar?: {
        width?: number;
        background?: string;
      };
      header?: {
        height?: number;
        background?: string;
      };
    }
  }
}

const theme = createTheme({
  palette: {
    primary: {
      main: '#1976d2',
    },
    secondary: {
      main: '#dc004e',
    },
  },
  custom: {
    sidebar: {
      width: 240,
      background: '#1a1a1a',
    },
    header: {
      height: 64,
      background: '#ffffff',
    },
  },
});

export default theme; 