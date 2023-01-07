import type { IGrade } from "./grade"

export interface IWeightedClass {
    classCode: string
    gradeList: Array<IGrade>
}