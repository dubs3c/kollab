

export interface Path {
    Id: string,
    Path: string,
    Verb: string,
    Headers: string[],
    Body: string
}


interface LogEventHeaders {
    [key: string]: string[]
}

export interface LogEvent {
    Id: string,
    Path: string,
    IP: string,
    RequestHeaders: LogEventHeaders
}