export interface ApiError {
    type: 
        // Public & Auth Errors
        | 'VALIDATION_ERROR' 
        | 'UNAUTHORIZED' 
        | 'INVALID_PASSWORD'
        | 'USER_NOT_FOUND'
        | 'DUPLICATE_USER'
        | 'CONFLICT_ERROR'
        
        // General API Errors
        | 'NOT_FOUND' 
        | 'RATE_LIMIT_EXCEEDED' 
        | 'INTERNAL_SERVER_ERROR' // Use this instead of SERVER_ERROR
        | 'NO_CONTENT' 
        | 'UNKNOWN_ERROR'
        
        // Domain Specific Errors
        | 'DUPLICATE_ITEM'; 
}