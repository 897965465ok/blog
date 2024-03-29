declare module '*.css';
declare module '*.less';
declare module '*.png';
declare module '*.scss' {
  const content: any;
  export default content;
}
declare module '*.svg' {
  export function ReactComponent(
    props: React.SVGProps<SVGSVGElement>,
  ): React.ReactElement;
  const url: string;
  export default url;
}
