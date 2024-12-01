export const ErrorCode = {
    SUCCESS: 200,
    PARAM_ERROR: 400,
    UNAUTHORIZED: 401,
    FORBIDDEN: 403,
    NOT_FOUND: 404,
    SERVER_ERROR: 500,
} as const

export const ErrorMessage = {
    [ErrorCode.SUCCESS]: '操作成功',
    [ErrorCode.PARAM_ERROR]: '参数错误',
    [ErrorCode.UNAUTHORIZED]: '未授权',
    [ErrorCode.FORBIDDEN]: '禁止访问',
    [ErrorCode.NOT_FOUND]: '资源不存在',
    [ErrorCode.SERVER_ERROR]: '服务器错误',
} as const 