// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {src} from '../models';
import {projectrelated} from '../models';

export function AddProjectToCollectionIfNeeded(arg1:string):Promise<void>;

export function AnyMadeProjectFilesInFolder(arg1:string):Promise<boolean>;

export function ChooseFolderForNewProject():Promise<string>;

export function ChooseProject():Promise<string>;

export function CreateProject(arg1:string,arg2:string,arg3:string,arg4:src.Loader):Promise<string>;

export function CurrentProject():Promise<projectrelated.MadeProject>;

export function CurrentProjectAddNewRecipe(arg1:string,arg2:{[key: string]: string}):Promise<projectrelated.HistoryItem>;

export function CurrentProjectChangeAction(arg1:string,arg2:string,arg3:string,arg4:{[key: string]: string}):Promise<void>;

export function CurrentProjectGetItemsTypeSuggestion(arg1:string):Promise<Array<string>>;

export function CurrentProjectHistory():Promise<Array<projectrelated.HistoryItem>>;

export function GetInformationToFillCreationForm(arg1:string):Promise<projectrelated.ProjectCreationInformation>;

export function GetProjects():Promise<Array<projectrelated.MadeProject>>;

export function OpenProjectInFileManager(arg1:string):Promise<string>;

export function SaveToFile():Promise<void>;

export function SetCurrentProjectByFolder(arg1:string):Promise<void>;
