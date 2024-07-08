"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __generator = (this && this.__generator) || function (thisArg, body) {
    var _ = { label: 0, sent: function() { if (t[0] & 1) throw t[1]; return t[1]; }, trys: [], ops: [] }, f, y, t, g;
    return g = { next: verb(0), "throw": verb(1), "return": verb(2) }, typeof Symbol === "function" && (g[Symbol.iterator] = function() { return this; }), g;
    function verb(n) { return function (v) { return step([n, v]); }; }
    function step(op) {
        if (f) throw new TypeError("Generator is already executing.");
        while (g && (g = 0, op[0] && (_ = 0)), _) try {
            if (f = 1, y && (t = op[0] & 2 ? y["return"] : op[0] ? y["throw"] || ((t = y["return"]) && t.call(y), 0) : y.next) && !(t = t.call(y, op[1])).done) return t;
            if (y = 0, t) op = [op[0] & 2, t.value];
            switch (op[0]) {
                case 0: case 1: t = op; break;
                case 4: _.label++; return { value: op[1], done: false };
                case 5: _.label++; y = op[1]; op = [0]; continue;
                case 7: op = _.ops.pop(); _.trys.pop(); continue;
                default:
                    if (!(t = _.trys, t = t.length > 0 && t[t.length - 1]) && (op[0] === 6 || op[0] === 2)) { _ = 0; continue; }
                    if (op[0] === 3 && (!t || (op[1] > t[0] && op[1] < t[3]))) { _.label = op[1]; break; }
                    if (op[0] === 6 && _.label < t[1]) { _.label = t[1]; t = op; break; }
                    if (t && _.label < t[2]) { _.label = t[2]; _.ops.push(op); break; }
                    if (t[2]) _.ops.pop();
                    _.trys.pop(); continue;
            }
            op = body.call(thisArg, _);
        } catch (e) { op = [6, e]; y = 0; } finally { f = t = 0; }
        if (op[0] & 5) throw op[1]; return { value: op[0] ? op[1] : void 0, done: true };
    }
};
var MoedaBrasil = {
    style: 'currency',
    currency: 'BRL',
    minimumFractionDigits: 2,
    maximumFractionDigits: 3
};
var formatarMoeda = new Intl.NumberFormat('pt-BR', MoedaBrasil);
var map = new Map();
for (var i = 0; i <= 200; i++) {
    var element = document.getElementById(i.toString());
    if (element) {
        console.log(element);
        map.set(i, true);
        continue;
    }
    map.set(i, false);
}
document.addEventListener("DOMContentLoaded", function () {
    var valores = document.querySelectorAll(".real");
    valores.forEach(function (valor) {
        if (valor.textContent) {
            valor.textContent = formatarMoeda.format(parseFloat(valor.textContent));
        }
    });
});
function scrollToId() {
    var Id = document.getElementById('searchPlu').value;
    if (Id !== null) {
        var element = document.getElementById(Id);
        if (element) {
            element.scrollIntoView({ behavior: 'smooth' });
        }
        else {
            alert("Produto não cadastrado");
        }
    }
}
function scrollToIdByValue(value) {
    var element = document.getElementById(value);
    if (element) {
        element.scrollIntoView({ behavior: 'smooth' });
    }
    else {
        console.log("Produto não cadastrado");
    }
}
function excluir(num) {
    var resposta = confirm("Certeza que deseja deletar o produto ".concat(num));
    if (!resposta) {
        return;
    }
    var bool = false;
    var linha = document.querySelector("#row-".concat(num));
    console.log(linha);
    if (!linha) {
        return;
    }
    linha.classList.add("fade-out");
    {
        var plu = parseInt(num);
        var dados = {
            Plu: plu
        };
        var json = JSON.stringify(dados);
        fetch('/excluir', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: json,
        })
            .then(function (response) {
            if (!response.ok) {
                return response.text().then(function (text) { throw new Error(text); });
            }
            return response.json();
        })
            .then(function (data) {
            console.log('Excluído:', data);
            // Recarregar a página após a exclusão bem-sucedida
            location.reload();
        })
            .catch(function (error) {
            try {
                var errorJson = JSON.parse(error.message);
                console.error('Erro ao excluir:', errorJson.message);
            }
            catch (e) {
                console.error('Erro ao excluir:', error.message);
            }
        });
    }
}
function toPageFile() {
    var resposta = confirm("\nA balan\u00E7a Clipse s\u00F3 cadastra produtos com c\u00F3digos de 1 a 200.\nOs demais produtos ser\u00E3o ignorados. \nDeseja continuar?");
    if (resposta) {
        window.location.href = "/file";
    }
}
function editar(num) {
    console.log(num);
    var inputName = ".input-plu-" + num;
    var inputs = document.querySelectorAll(inputName);
    for (var i = 0; i < inputs.length; i++) {
        var inp = inputs[i];
        if (inp.style.display == "none") {
            inp.style.display = "inline";
            continue;
        }
        inp.style.display = "none";
    }
}
function exibeNew() {
    var element = document.getElementById("newPlu");
    if (!element) {
        return;
    }
    console.log(element.style.display);
    if (element.style.display == "none") {
        element.style.display = "";
        return;
    }
    element.style.display = "none";
}
function enviarDados(dados) {
    var json = JSON.stringify(dados);
    console.log(json);
    return fetch('/editar', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: json,
    })
        .then(function (response) {
        if (!response.ok) {
            return response.text().then(function (text) { throw new Error(text); });
        }
        return response.json();
    })
        .then(function (data) {
        console.log('Sucesso:', data);
        location.reload();
        return data;
    })
        .catch(function (error) {
        try {
            var errorJson = JSON.parse(error.message);
            console.error('Erro:', errorJson.message);
        }
        catch (e) {
            console.error('Erro:', error.message);
        }
        throw error; // Rejeitar a promessa com o erro
    });
}
function coletarDados(plu) {
    return new Promise(function (resolve, reject) {
        var descricaoElement = document.querySelector("input[name=\"descricao-".concat(plu, "\"]"));
        var precoElement = document.querySelector("input[name=\"preco-".concat(plu, "\"]"));
        var vendaElement = document.querySelector("select[name=\"venda-".concat(plu, "\"]"));
        var validadeElement = document.querySelector("input[name=\"validade-".concat(plu, "\"]"));
        if (!descricaoElement || !precoElement || !vendaElement || !validadeElement) {
            console.error('Elementos de dados não encontrados.');
            alert('Erro ao coletar dados.');
            return reject(new Error('Elementos de dados não encontrados.'));
        }
        var descricao = descricaoElement.value.trim().toUpperCase();
        var preco = parseFloat(precoElement.value.trim());
        var vendaSelect = parseInt(vendaElement.value.trim(), 10);
        var validade = parseInt(validadeElement.value.trim(), 10);
        var bolean = descricao === '' || isNaN(preco) || isNaN(vendaSelect) || isNaN(validade);
        if (bolean) {
            console.log('Erro no cadastro');
            alert('Erro no cadastro');
            return reject(new Error('Erro no cadastro'));
        }
        var dados = {
            Plu: parseInt(plu, 10),
            Descricao: descricao,
            Preco: preco,
            Venda: vendaSelect,
            Validade: validade,
        };
        console.log(dados);
        resolve(dados);
    });
}
function coletarTodosDados() {
    var linhas = document.querySelectorAll('tbody tr');
    var promessasDados = [];
    linhas.forEach(function (linha) {
        var _a;
        var pluElement = linha.querySelector('td[id]');
        if (pluElement) {
            var texto = ((_a = pluElement.textContent) !== null && _a !== void 0 ? _a : "").trim();
            var dadosPLU = coletarDados(texto);
            promessasDados.push(dadosPLU);
        }
    });
    return Promise.all(promessasDados)
        .then(function (dadosProdutos) {
        return enviarDados(dadosProdutos);
    })
        .then(function () {
        console.log('Todos os dados foram enviados com sucesso.');
    })
        .catch(function (error) {
        console.error('Erro ao enviar dados:', error);
        alert('Erro ao enviar dados.');
    });
}
function novo() {
    return new Promise(function (resolve, reject) {
        var pluHtml = document.getElementById("plu-new");
        if (!pluHtml) {
            return resolve(false);
        }
        var descricaoElement = document.getElementById("descricao-new");
        var precoElement = document.getElementById("preco-new");
        var vendaSelectElement = document.querySelector("select[name=\"venda-new\"]");
        var validadeElement = document.getElementById("validade-new");
        if (!descricaoElement || !precoElement || !vendaSelectElement || !validadeElement) {
            return resolve(false);
        }
        var descricao = descricaoElement.value.toUpperCase();
        var preco = parseFloat(precoElement.value);
        var vendaSelect = parseInt(vendaSelectElement.value.trim(), 10);
        var validade = parseInt(validadeElement.value, 10);
        console.log(pluHtml.value);
        var pluInvalido = isNaN(parseInt(pluHtml.value)) || descricao === "" || isNaN(preco) || isNaN(vendaSelect) || isNaN(validade);
        if (pluInvalido) {
            alert("Novo cadastro inválido");
            return resolve(false);
        }
        var plu = parseInt(pluHtml.value, 10);
        if (plu > 200) {
            alert("Produto só pode ser cadastrado do código 1 ao 200");
            return resolve(false);
        }
        if (validade > 200) {
            alert("Validade do produto tem limite de até 200 dias");
            return resolve(false);
        }
        var dados = {
            Plu: plu,
            Descricao: descricao,
            Preco: preco,
            Venda: vendaSelect,
            Validade: validade,
        };
        enviarDadosNovoPlu(dados)
            .then(function () { return resolve(true); })
            .catch(function () { return resolve(false); });
    });
}
function enviarDadosNovoPlu(dados) {
    var json = JSON.stringify(dados);
    console.log(json);
    return fetch('/novo', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: json,
    })
        .then(function (response) {
        if (!response.ok) {
            return response.text().then(function (text) { throw new Error(text); });
        }
        return response.json();
    })
        .then(function (data) {
        console.log('Sucesso:', data);
        location.reload();
    })
        .catch(function (error) {
        try {
            var errorJson = JSON.parse(error.message);
            console.error('Erro:', errorJson.message);
        }
        catch (e) {
            console.error('Erro:', error.message);
        }
        throw error;
    });
}
function startImport(event) {
    return __awaiter(this, void 0, void 0, function () {
        var form, campos, todosPreenchidos, aviso, progressElement, statusElement, ellipsisElement, progressBar, ellipsisCounter, ellipsisInterval, pulsateInterval, formulario, formData, request;
        return __generator(this, function (_a) {
            reiniciarCronometro();
            iniciarCronometro();
            event.preventDefault(); // Prevent default form submission
            form = document.getElementById('importForm');
            if (!form) {
                return [2 /*return*/];
            }
            campos = form.querySelectorAll('input[required]');
            todosPreenchidos = true;
            campos.forEach(function (campo) {
                if (!campo.value.trim()) {
                    todosPreenchidos = false;
                }
            });
            if (!todosPreenchidos) {
                aviso = 'Por favor, preencha todos os campos obrigatórios.';
                alert(aviso);
                return [2 /*return*/];
            }
            progressElement = document.getElementById("progress");
            statusElement = document.getElementById("status");
            ellipsisElement = document.getElementById("ellipsis");
            progressBar = document.querySelector(".progress-bar");
            if (progressElement) {
                // Reset progress bar and status
                progressElement.style.width = '0%';
                progressElement.textContent = '0%';
            }
            if (statusElement) {
                statusElement.textContent = 'Importando...';
            }
            if (progressBar) {
                progressBar.style.display = 'block';
            }
            ellipsisCounter = 0;
            ellipsisInterval = setInterval(function () {
                if (!ellipsisElement) {
                    return;
                }
                ellipsisElement.textContent = '.'.repeat((ellipsisCounter % 4) + 1);
                ellipsisCounter++;
            }, 500);
            pulsateInterval = setInterval(function () {
                if (!statusElement) {
                    return;
                }
                statusElement.style.fontSize = '1.01em';
                setTimeout(function () {
                    if (!statusElement) {
                        return;
                    }
                    statusElement.style.fontSize = '1em';
                }, 250);
            }, 500);
            formulario = document.getElementById("importForm");
            if (!formulario) {
                return [2 /*return*/];
            }
            formData = new FormData(formulario);
            request = new XMLHttpRequest();
            request.open("POST", "/importarteste", true);
            request.setRequestHeader("Accept", "text/event-stream");
            request.onreadystatechange = function () {
                if (request.readyState === 3 || request.readyState === 4) {
                    var lines = request.responseText.split("\n\n");
                    lines.forEach(function (line) {
                        if (line.startsWith("data: ")) {
                            var data = line.replace("data: ", "").trim();
                            if (data) {
                                try {
                                    var jsonData = JSON.parse(data);
                                    if (jsonData.hasOwnProperty("progress")) {
                                        var progressValue = jsonData.progress;
                                        if (!progressElement) {
                                            return;
                                        }
                                        progressElement.style.width = progressValue + '%';
                                        progressElement.textContent = Math.round(progressValue) + '%';
                                    }
                                    if (jsonData.hasOwnProperty("complete")) {
                                        clearInterval(ellipsisInterval); // Stop ellipsis animation
                                        clearInterval(pulsateInterval); // Stop pulsate animation
                                        if (!ellipsisElement || !statusElement) {
                                            return;
                                        }
                                        ellipsisElement.textContent = ''; // Clear ellipsis
                                        statusElement.textContent = 'Finalizado';
                                        reiniciarCronometro();
                                    }
                                }
                                catch (e) {
                                    // Handle non-JSON data
                                    console.error("Error parsing JSON data:", e);
                                }
                            }
                        }
                    });
                }
                if (request.readyState === 4 && request.status !== 200) {
                    if (ellipsisElement && statusElement) {
                        statusElement.textContent = 'Erro ao completar a tarefa.';
                        ellipsisElement.textContent = ''; // Clear ellipsis
                    }
                    reiniciarCronometro();
                    clearInterval(ellipsisInterval); // Stop ellipsis animation
                    clearInterval(pulsateInterval); // Stop pulsate animation
                }
            };
            request.send(formData);
            return [2 /*return*/];
        });
    });
}
var importForm = document.getElementById("importForm");
if (importForm) {
    importForm.addEventListener("submit", startImport);
}
function formatarData(eventDate) {
    var parts = eventDate.split(" ");
    // Divide a string na primeira ocorrência de espaço
    var complete_time = parts[1];
    var data_particionada = parts[0].split("-");
    var ano = data_particionada[0];
    var mes = data_particionada[1];
    var dia = data_particionada[2];
    var data = "".concat(dia, "/").concat(mes, "/").concat(ano);
    var time = complete_time.split(".")[0];
    // Agora parts[0] contém "2024-06-25"
    var dataFormatada = "Dia: ".concat(data, " \u00E0s: ").concat(time);
    return dataFormatada;
}
var tempoSegundos = 0;
var cronometroAtivo = false;
var intervalo;
function formatarTempo(segundos) {
    var horas = String(Math.floor(segundos / 3600)).padStart(2, '0');
    var minutos = String(Math.floor((segundos % 3600) / 60)).padStart(2, '0');
    var seg = String(segundos % 60).padStart(2, '0');
    return "".concat(horas, ":").concat(minutos, ":").concat(seg);
}
function atualizarDisplay() {
    document.getElementById('tempo').innerText = formatarTempo(tempoSegundos);
}
function iniciarCronometro() {
    return __awaiter(this, void 0, void 0, function () {
        return __generator(this, function (_a) {
            if (cronometroAtivo)
                return [2 /*return*/];
            cronometroAtivo = true;
            intervalo = setInterval(function () {
                tempoSegundos++;
                atualizarDisplay();
            }, 1000);
            return [2 /*return*/];
        });
    });
}
function pausarCronometro() {
    cronometroAtivo = false;
    clearInterval(intervalo);
}
function reiniciarCronometro() {
    pausarCronometro();
    tempoSegundos = 0;
    atualizarDisplay();
}
var campoFiltro = document.getElementById("searchPluDescription");
if (campoFiltro) {
    campoFiltro.addEventListener("input", function () {
        var _this = this;
        console.log(this.value);
        var produtos = document.querySelectorAll(".descricao-todos-produtos");
        produtos.forEach(function (produto) {
            var expressao = new RegExp(_this.value, "i");
            var id = produto.id;
            var descricao_codigo = id.split("-");
            var codigo = descricao_codigo[1];
            var descricao = produto.textContent;
            var idLinhaTabela = "row-".concat(codigo);
            var linha = document.getElementById(idLinhaTabela);
            if (linha && descricao) {
                if (!expressao.test(descricao)) {
                    linha.classList.add("invisivel");
                }
                else {
                    linha.classList.remove("invisivel");
                }
            }
        });
    });
}
//# sourceMappingURL=scripts.js.map