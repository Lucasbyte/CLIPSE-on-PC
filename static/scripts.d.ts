declare const MoedaBrasil: Intl.NumberFormatOptions;
declare const formatarMoeda: Intl.NumberFormat;
declare var map: any;
declare function scrollToId(): void;
declare function scrollToIdByValue(value: string): void;
declare function excluir(num: string): void;
declare function toPageFile(): void;
declare function editar(num: string): void;
declare function exibeNew(): void;
interface DadosProduto {
    Plu: number;
    Descricao: string;
    Preco: number;
    Venda: number;
    Validade: number;
}
interface RespostaServidor {
    message: string;
}
declare function enviarDados(dados: DadosProduto[]): Promise<RespostaServidor[]>;
declare function coletarDados(plu: string): Promise<DadosProduto>;
declare function coletarTodosDados(): Promise<void>;
declare function novo(): Promise<boolean>;
declare function enviarDadosNovoPlu(dados: DadosProduto): Promise<void>;
declare function startImport(event: Event): Promise<void>;
declare let importForm: HTMLElement | null;
declare function formatarData(eventDate: string): string;
declare let tempoSegundos: number;
declare let cronometroAtivo: boolean;
declare let intervalo: number;
declare function formatarTempo(segundos: number): string;
declare function atualizarDisplay(): void;
declare function iniciarCronometro(): Promise<void>;
declare function pausarCronometro(): void;
declare function reiniciarCronometro(): void;
declare var campoFiltro: HTMLInputElement;
